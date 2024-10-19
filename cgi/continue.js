/**
 * This script continues a payment that has been accepted or rejected
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

  let finalizedOutgoingPaymentGrant;

  const grantContinuationErrorMessage =
    "There was an error continuing the grant. You probably have not accepted the grant at the url (or it has already been used up, in which case, rerun the script).";

  try {
    finalizedOutgoingPaymentGrant = await client.grant.continue({
      url: process.env.APP_IH2024_CONTINUE_URI,
      accessToken: process.env.APP_IH2024_CONTINUE_ACCESS_TOKEN,
    });
  } catch (err) {
    if (err instanceof OpenPaymentsClientError) {
      console.log({
        "Message": grantContinuationErrorMessage,
      });
      process.exit(1);
    }

    throw err;
  }

  if (!isFinalizedGrant(finalizedOutgoingPaymentGrant)) {
    console.log({
      "Message": "There was an error continuing the grant. You probably have not accepted the grant at the url.",
    });
    process.exit(1);
  }

  // console.log(
  //   "\nStep 6: got finalized outgoing payment grant",
  //   finalizedOutgoingPaymentGrant
  // );

  // Step 7: Finally, create the outgoing payment on the sending wallet address.
  // This will make a payment from the outgoing payment to the incoming one (over ILP)
  const outgoingPayment = await client.outgoingPayment.create(
    {
      url: sendingWalletAddress.resourceServer,
      accessToken: finalizedOutgoingPaymentGrant.access_token.value,
    },
    {
      walletAddress: sendingWalletAddress.id,
      quoteId: process.env.APP_IH2024_QUOTE_ID,
    }
  );

  // console.log(
  //   "\nStep 7: Created outgoing payment. Funds will now move from the outgoing payment to the incoming payment.",
  //   outgoingPayment
  // );
  console.log({
    "Message": "Created outgoing payment",
  });
  process.exit(0);
})();
