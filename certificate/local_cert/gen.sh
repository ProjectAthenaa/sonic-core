rm *.pem

openssl req -nodes -x509 -newkey rsa:4096 -days 365 -keyout ca-key.pem -out ca-cert.pem -subj "/C=US/ST=NewYork/L=NewYorkCity/O=ProjectAthena/OU=Software/CN=*.localhost/emailAddress=info@athenabot.com"

echo "CA's self-signed certificate"

openssl x509 -in ca-cert.pem -noout -text

openssl req -nodes -newkey rsa:4096 -keyout server-key.pem -out server-req.pem -subj "/C=US/ST=NewYork/L=NewYorkCity/O=ProjectAthena/OU=Software/CN=*.localhost/emailAddress=info@athenabot.com"

openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem

echo "Server's signed certificate"

openssl x509 -in server-cert.pem -noout -text

openssl req -nodes -newkey rsa:4096 -keyout client-key.pem -out client-req.pem -subj "/C=US/ST=NewYork/L=NewYorkCity/O=ProjectAthena/OU=Software/CN=*.localhost/emailAddress=info@athenabot.com"

openssl x509 -req -in client-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem

echo "Client's signed certificate"
openssl x509 -in client-cert.pem -noout -text
