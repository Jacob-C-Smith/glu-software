

printf "
password\n
password\n
US\n
Idaho\n
Boise\n
Glu Software\n
\n
Jacob Smith\n
jacobsmithsinbox@gmail.com\n
" |
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes