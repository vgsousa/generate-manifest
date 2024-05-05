# Aplicação Go para Configuração de Deployment Kubernetes
Este é um exemplo de aplicação em Go que recebe requisições HTTP para criar configurações de deployments em um cluster Kubernetes. A aplicação é capaz de gerar configurações YAML, como Ingress, a partir de dados fornecidos em requisições JSON.

## Funcionalidades
- Recebe requisições HTTP contendo dados de aplicação em formato JSON.
- Valida se todos os campos obrigatórios foram preenchidos no JSON.
- Gera arquivos de configuração YAML baseados em templates.
- Permite o download dos arquivos YAML gerados.

## Pré-requisitos
Certifique-se de ter o seguinte instalado em sua máquina local:

- Go (você pode baixá-lo em golang.org)
- Docker (você pode baixá-lo em docker.com)

## Instalação e Uso
1.Clone este repositório em sua máquina local:
```bash
git clone https://github.com/vgsousa/generate-manifest.git
```

2.Navegue até o diretório do projeto:
```bash
cd generate-manifest
```

3.Construa a imagem Docker:
```bash
docker build . -t generate-manifest
```

4.Execute o contêiner Docker:
```bash
docker run -p 8080:8080 generate-manifest
```
5.Sua aplicação estará disponível em http://localhost:8080.



# API
## 'POST /create-deployment'
Este endpoint recebe uma requisição contendo os dados da aplicação em formato JSON. Os campos obrigatórios são:

- project: Nome da aplicação.
- team: Time responsável pela aplicação.
- environment: Ambiente da aplicação.
- namespace: Namespace da aplicação.
- port: Porta principal da aplicação.

Exemplo de requisição:

```json
{
    "project": "aplication-name",
    "team": "DSCOVR",
    "environment": "dev",
    "namespace": "web",
    "port": 80
}
```
A resposta da API será um arquivo YAML contendo as configurações geradas para o Kubernetes.

Licença
Este projeto está licenciado sob a MIT License.