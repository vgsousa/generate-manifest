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
## 'POST /generate'
Este endpoint recebe uma requisição contendo os dados da aplicação em formato JSON. Os campos obrigatórios são:

- application: Nome da aplicação.
- clusterVersion: Versão do Cluster.
- environment: Ambiente da aplicação.
- internalAccessOnly: Bloqueio de acesso por IP.
- registry: Registry das imagens.
- namespace: Namespace da aplicação.
- replicas: Número de pods.
- port: Porta principal da aplicação.
- resources: Recursos requeridos e limitado.
- livenessProbePath: URL do LivenessProbe.
- readinessProbePath: URL do ReadinessProbe.
- hpa: Configurações de mínimo, máximo e quais as métricas.

Exemplo de requisição:

```json
{
    "application": "Microservices api",
    "environment": "dev",
    "clusterVersion": 1.99,
    "internalAccessOnly": true,
    "registry": "registry.azurecr.io",
    "namespace": "portal",
    "deployment": {
        "replicas": 1,
        "port": 80,
        "resources": {
            "requests" :{
                "memory": "350Mi",
                "cpu": "250m"
            },
            "limits" :{
                "memory": "1Gi",
                "cpu": "512m"
            }
        },
        "livenessProbePath": "/health-check",
        "readinessProbePath": "/health-check"
    },
    "hpa" :{
        "min": 2,
        "max": 10,
        "metrics": {
            "cpu": 50,
            "memory": 80
        }
    }
}
```
A resposta da API será um arquivo YAML contendo as configurações geradas para o Kubernetes.

Licença
Este projeto está licenciado sob a MIT License.