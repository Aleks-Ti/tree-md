# База

У раста есть единый rustup которы и раст обновит и карго(пакетный мендежер) поставит.

## rustup

For linux

```bash
$ curl --proto '=https' --tlsv1.2 https://sh.rustup.rs -sSf | sh
```

or windows

Перейдите на сайт `https://www.rust-lang.org/tools/install`, скачайте и запустите установщик rustup-init.exe.

Предложит на выбор. Выбираем быструю установку `1`

```bash
Rust Visual C++ prerequisites

Rust requires a linker and Windows API libraries but they don't seem to be
available.

These components can be acquired through a Visual Studio installer.

1) Quick install via the Visual Studio Community installer
   (free for individuals, academic uses, and open source).

2) Manually install the prerequisites
   (for enterprise and advanced users).

3) Don't install the prerequisites
   (if you're targeting the GNU ABI).

>
```

Проверка версии:

```bash
rustc --version
```

```bash
rustup update
```

## cargo

Создать проект:

```bash
cargo new hello_world
```

Создать библиотеку:

```bash
cargo new hello_world --lib
```

### Base command for cargo

```bash
# список команд
cargo --list
```

```bash
# скомпилировать и запустить проект
cargo run
```

```bash
# скомпилировать
cargo build
```

```bash
# проверить на ошибки
cargo check
```

### форматер rustfmt

добавить форматер

```bash
rustup component add rustfmt
```

После, в директории проекта можно использовать команду для форматирования:

```bash
cargo fmt
```

По умолчанию `rustfmt` использует стандартные правила. Если по каким-то причинам их понадобится кастомизировать, в корне проекта можно создать файл `rustfmt.toml` или .`rustfmt.toml` с примерным содержимым:

```toml
max_width = 80    # <- максимальная длина строки
tab_spaces = 4    # <- количество пробелов отступа
```

Полный список команд -> `https://rust-lang.github.io/rustfmt.`

### линтер clippy

Установить:

```bash
rustup component add clippy
```

Находясь в директории проекта, запустите команду:

```bash
cargo clippy
```
