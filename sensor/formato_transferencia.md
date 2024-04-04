# Formato dos dados de transferência entre o broker e os sensores

## Broker -> Sensores

```text
Id: <id>
Cmd: <comando>

<conteúdo (opcional)>
```

## Sensores -> Broker (retorno do comando)

```text
Id: <id>
Cmd: <comando>

<conteúdo (opcional)>
```