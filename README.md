# funcgoaws
Projeto para estudar integração AWS Lambda 
com Golang usando banco DynamoDB.

# Necessidades de ambiente
- AWS Enviroment, com credenciais configuradas pra uso;
- Lambda function criada pra runtime Go 1.x, com Role de leitura ao DynamoDB;
- Table chamada Toggle com uma Partition Key definida como "toggleName" e dois outros campos "descrição" e "enable"; 
- AWS CLI e AWS SAM caso queira mais praticidade pra testes locais direto na AWS;
- Go SDK instalado;

Para gerar o pacote para subir pra Lambda,use os comandos:

* Empacotar o fonte a partir da classe principal

``` GOOS=linux go build main.go ```
* Compactar o arquivo para um tipo válido (zip file) para a Lambda

``` zip function.zip main ```

Após isso, na raiz do projeto o .zip estará disponível
para subir para a Lambda.

Para invocar a Lambda via AWS console, ou AWS CLI ou até AWS SAM.

Você invoca passando um objeto do tipo APIGatewayRequestProxy

Método disponível: GET

URL:/Toggle

QueryParam: toggleName

Retorno: objeto do tipo APIGatewayResponseProxy 
