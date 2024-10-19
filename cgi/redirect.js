/**
 * This script sets up an incoming payment on a receiving wallet address,
 * and a quote on the sending wallet address (after getting grants for both of the resources).
 *
 * The final step is asking for an outgoing payment grant for the sending wallet address.
 * The user will need to navigate to the URL, and accept, or rejecct, the interactive grant.
 */

import {
  createAuthenticatedClient,
  OpenPaymentsClientError,
  isFinalizedGrant,
} from "@interledger/open-payments";

(async () => {
  const client = await createAuthenticatedClient({
    walletAddressUrl: process.env.APP_IH2024_IN_WALLET_ADDRESS_URL,
    keyId: process.env.APP_IH2024_KEY_ID,
    privateKey: process.env.APP_IH2024_PRIVATE_KEY,
  });

  const sendingWalletAddress = await client.walletAddress.get({
    url: process.env.APP_IH2024_IN_WALLET_ADDRESS_URL,
  });
  const receivingWalletAddress = await client.walletAddress.get({
    url: process.env.APP_IH2024_OUT_WALLET_ADDRESS_URL,
  });

  // console.log(
  //   "Got wallet addresses. We will set up a payment between the sending and the receiving wallet address",
  //   { receivingWalletAddress, sendingWalletAddress }
  // );

  // Step 1: Get a grant for the incoming payment, so we can create the incoming payment on the receiving wallet address
  const incomingPaymentGrant = await client.grant.request(
    {
      url: receivingWalletAddress.authServer,
    },
    {
      access_token: {
        access: [
          {
            type: "incoming-payment",
            actions: ["read", "complete", "create"],
          },
        ],
      },
    }
  );

  // console.log(
  //   "\nStep 1: got incoming payment grant for receiving wallet address",
  //   incomingPaymentGrant
  // );

  // Step 2: Create the incoming payment. This will be where funds will be received.
  const incomingPayment = await client.incomingPayment.create(
    {
      url: receivingWalletAddress.resourceServer,
      accessToken: incomingPaymentGrant.access_token.value,
    },
    {
      walletAddress: receivingWalletAddress.id,
      incomingAmount: {
        assetCode: receivingWalletAddress.assetCode,
        assetScale: receivingWalletAddress.assetScale,
        value: process.env.APP_IH2024_AMOUNT,
      },
    }
  );

  // console.log(
  //   "\nStep 2: created incoming payment on receiving wallet address",
  //   incomingPayment
  // );

  // Step 3: Get a quote grant, so we can create a quote on the sending wallet address
  const quoteGrant = await client.grant.request(
    {
      url: sendingWalletAddress.authServer,
    },
    {
      access_token: {
        access: [
          {
            type: "quote",
            actions: ["create", "read"],
          },
        ],
      },
    }
  );

  // console.log(
  //   "\nStep 3: got quote grant on sending wallet address",
  //   quoteGrant
  // );

  // Step 4: Create a quote, this gives an indication of how much it will cost to pay into the incoming payment
  const quote = await client.quote.create(
    {
      url: sendingWalletAddress.resourceServer,
      accessToken: quoteGrant.access_token.value,
    },
    {
      walletAddress: sendingWalletAddress.id,
      receiver: incomingPayment.id,
      method: "ilp",
    }
  );

  // console.log("\nStep 4: got quote on sending wallet address", quote);

  // Step 5: Start the grant process for the outgoing payments.
  // This is an interactive grant: the user (in this case, you) will need to accept the grant by navigating to the outputted link.
  let interact = {
    start: ["redirect"],
  }
  if (process.env.APP_IH2024_SUCCESS_URL != "") {
    interact.finish = {
      method: "redirect",
      uri: process.env.APP_IH2024_SUCCESS_URL,
      nonce: process.env.APP_IH2024_NONCE,
    }
  }
  const outgoingPaymentGrant = await client.grant.request(
    {
      url: sendingWalletAddress.authServer,
    },
    {
      access_token: {
        access: [
          {
            type: "outgoing-payment",
            actions: ["create"],
            limits: {
              debitAmount: quote.debitAmount,
            },
            identifier: sendingWalletAddress.id,
          },
        ],
      },
      interact: interact,
    }
  );

  // console.log(
  //   "\nStep 5: got pending outgoing payment grant",
  //   outgoingPaymentGrant
  // );
  // console.log(
  //   "Please navigate to the following URL, to accept the interaction from the sending wallet:"
  // );
  console.log({
    "Redirect": outgoingPaymentGrant.interact.redirect,
    "ContinueURI": outgoingPaymentGrant.continue.uri,
    "AccessToken": outgoingPaymentGrant.continue.access_token.value,
    "QuoteID": quote.id,
  });

  process.exit(0);
})();
