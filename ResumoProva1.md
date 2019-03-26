# Resumo Prova 1 Redes

## What Is the Internet ?
Definir o que é internet não é um processo muito simples, muito de nós simplesmente pensam que "Internet" é aquilo que aparece quando clicamos no navegador e abre uma página de pesquisa,  podemos descrever a Internet basicamente em termos de uma rede de infraestrutura que fornece serviços para aplicações distribuídas.

A internet é uma rede de computadores que interconecta bilhões de dispositivos pelo mundo. Na internet os dispositivos são chamados de ***hosts*** ou ***end systems***.

 ***End systems :***
- São conectados por um ***link / enlace*** e ***comutadores de pactes***
- Quando um end system tem dados pra mandar pra outro end system, o que sistema que irá mandar os dados segmenta ele e add um header (em bytes) pra cada segmento, o resultado final disso se chama ***packets***, que são os pacotes de dados que irá ser enviado pela rede.
- Um comutador de pacotes pega esse pacote chegando de algum link e passa esse pacote para frente.
	- Um comutador de pacote tem vários formatos e tipos, os 2 mais famosos são os ***roteadores*** e ***switches*** (aquele pra dividir a net pra jogar com os amigos em casa) 

***Comutadores de Pacotes :***
- Os 2 são usados para passar os pacotes para frente. ***switches*** são mais usandos em redes de acesso, enquanto os ***roteadores*** são mais usados na ***rede de núcleo***.
- Essa sequencia de links e comutadores de pacotes junto com os pacotes do sistema final que envia até o sistema final que recebe é conhecido como ***route*** ou ***path***.
- ***Obs:*** Uma analogia a se fazer é a seguinte: Uma fabrica precisa mandar um monte de material para um destinatário. Na fábrica, esse material é dividido e colocados em vários caminhões, os caminhões vão (independente) bonitinhos pela pista, de vez em quando passam por algumas interseções (pedágios ?) até o destinatário, chegando lá os caminhões vão descarregando o material e juntos eles de novo formando o que tinha antes. Isso pode ser comparado com :
	- Fábrica : End System que manda os dados
	- Caminhão : Pacotes
	- Material dentro do caminhão : os bits (dado)
	- Pista : Links
	- Pedágio : Comutadores de pacotes
	- Destinatário : Sistema final que recebe os dados
- Os sistemas finais acessam a Internet através da ***Internet Service Providers (ISPs)***, incluindo ISPs residencial, como cabos locais, companhias de telefone (tem uma porrada de exemplo). Cada ISP é por si só uma rede de comutação de pacotes e links, ele provide [buguei pra traduzir] vários tipos de rede de acesso para os sistemas finais, incluindo acesso a banda laga como cabo do modem, mobile wireless access e essas porras.

[TEM MAIS COISAS]

### Serviços

## The Network Edge

## The Network Core

## Delay, Loss, and Throughput in Packet-Switched Networks

## Protocol Layers and Their Service Models



