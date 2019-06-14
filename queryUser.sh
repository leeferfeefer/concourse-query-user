#!/bin/sh

userinput1=$(osascript -e 'tell app "System Events" to set LDAP1 to the text returned of (display dialog "What is user 1 LDAP?" default answer "")')
passinput1=$(osascript -e 'tell app "System Events" to set PASS1 to the text returned of (display dialog "What is user 1 password?" with hidden answer default answer "")')
userinput2=$(osascript -e 'tell app "System Events" to set LDAP2 to the text returned of (display dialog "What is user 2 LDAP?" default answer "")')
passinput2=$(osascript -e 'tell app "System Events" to set PASS2 to the text returned of (display dialog "What is user 2 password?" with hidden answer default answer "")')
echo "$userinput1:$userinput2"
