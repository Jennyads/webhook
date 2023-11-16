<h2>Anotações</h2>
webhook - forma com que dois sistemas na internet conseguem se comunicar (enviar e receber webhook).

Geralmente é enviado por JSON. É requisitado por HTTP.
GET -> receber webhook
POST -> enviar webhook

WEBHOOK - requisições/comunicações (troca de informações quando uma ação acontece)
JSON - formato em que a comunicação acontece
API - sistema inteiro 


<h2>Definições</h2>
Um webhook é uma forma de comunicação entre sistemas, onde um sistema pode enviar automaticamente informações para outro assim que algo acontece. Geralmente, isso envolve a definição de um ponto de extremidade (URL) no sistema receptor, e quando um evento específico ocorre no sistema emissor, ele envia uma solicitação HTTP para o URL configurado.

A solicitação HTTP enviada pelo sistema emissor pode conter dados relevantes sobre o evento que ocorreu. O sistema receptor, que possui o webhook configurado, pode então processar esses dados de acordo com as suas necessidades.


Exemplo em python:
```
import requests
import json

#envia um webhook
link = "https://webhook.site/3ae3e0a2-ff88-42e0-a713-b5f2a70b799d"

dicionario = {
"nome": "Jenny",
"valor": 999,
"parcelas": 12,
}

dicionario = json.dumps(dicionario)
requests.post(link, data=dicionario_ajustado)


#endpoint de uma api
#flask
@ -> libera esse link de receber um POST
def receber_webhook():
    dicionario = request.json
    #seguir fazer o que o código quer 

```
