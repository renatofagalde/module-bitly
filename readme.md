# 🧩 module-bitly

### 🔢 Conversão Base62 para IDs curtos, rápidos e amigáveis (tipo Bitly)

![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go)
![License](https://img.shields.io/badge/license-MIT-green)
![Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)

O **module-bitly** é um módulo simples, rápido e direto para conversão:

-   🔷 `uint64` → Base62\
-   🔷 Base62 → `uint64`\
-   🔶 `[]byte` (ex.: UUID) → Base62\
-   🔶 Base62 → `[]byte`

Ideal para gerar **IDs curtos**, amigáveis, comparáveis e seguros para
URLs --- no estilo **Bitly** ou **YouTube ID**.

------------------------------------------------------------------------

# ✨ Recursos

-   🚀 **Super leve** --- zero dependências externas\
-   ⚡ **Rápido** --- implementado com operações matemáticas de baixo
    nível\
-   🔁 **Encode/Decode reversível**\
-   🧬 **Suporte a UUID em bytes**\
-   🛡 **Validação de caracteres e overflow**\
-   🧪 **100% coberto por testes unitários**\
-   📦 **Pronto para ser importado via Go Modules**

------------------------------------------------------------------------

# 📦 Instalação

    go get github.com/renatofagalde/module-bitly

------------------------------------------------------------------------

# 🛠 Uso

## 🔷 1. Converter `uint64` → Base62

``` go
import "github.com/renatofagalde/module-bitly"

id := uint64(123456789)
short := bitly.E(id)

fmt.Println(short)  // "8M0kX"
```

------------------------------------------------------------------------

## 🔷 2. Converter Base62 → `uint64`

``` go
decoded, err := bitly.D("8M0kX")
if err != nil {
    panic(err)
}

fmt.Println(decoded) // 123456789
```

------------------------------------------------------------------------

## 🔶 3. Converter UUID (16 bytes) → Base62

``` go
u := uuid.New()       // github.com/google/uuid
short := bitly.EncodeBytes(u[:])

fmt.Println(short) // exemplo: "5B3cf29AMQbF2xE8c"
```

------------------------------------------------------------------------

## 🔶 4. Base62 → UUID bytes

``` go
bytesUUID, err := bitly.DecodeBytes(short)
if err != nil {
    panic(err)
}
fmt.Printf("%x
", bytesUUID)
```

------------------------------------------------------------------------

# 🧪 Testes

Rodar testes com cobertura:

    go test ./... -cover

Gerar relatório HTML:

    go test ./... -coverprofile=coverage.out
    go tool cover -html=coverage.out -o coverage.html

------------------------------------------------------------------------

# 💡 Por que Base62?

-   URLs amigáveis\
-   IDs mais curtos\
-   Menos erros humanos ao copiar/colar\
-   Evita caracteres especiais de URL\
-   Excelente para encurtar UUIDs

------------------------------------------------------------------------

# 🧱 Estrutura Interna

    module-bitly/
    ├── bitly.go        # Implementação Base62
    ├── bitly_test.go   # Testes completos e 100% de cobertura
    └── go.mod

------------------------------------------------------------------------

# ⚠️ Limitações

-   Apenas caracteres válidos (`0–9`, `A–Z`, `a–z`) são aceitos\
-   Para valores acima de `uint64`, use `EncodeBytes` com arrays de
    bytes\
-   Ordem lexicográfica não é igual à ordem numérica (isso não é
    Base36/Base62 ordenável)

------------------------------------------------------------------------

# ❤️ Contribuição

PRs e sugestões são bem-vindas!

------------------------------------------------------------------------

# 📄 Licença

MIT License.
