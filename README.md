# lex-lambda-sample-init
Lex Lambda Sample Init


# To test locally
sam build
sam local generate-event connect contact-flow-event | sam local invoke -e -


# To package and deploy
make deployv2