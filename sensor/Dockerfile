# create a Dockerfile for python
FROM python:3.11.5

# Define o diretório de trabalho
WORKDIR /app

# Copia os arquivos .py para o diretório de trabalho
COPY . .

EXPOSE 3333

# Define o comando de execução
CMD ["python", "main.py"]