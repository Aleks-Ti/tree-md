# Инструкция по использованию горячих клавиш в редакторе VI/Vim

Редактор VI/Vim имеет несколько режимов работы и множество команд, которые позволяют эффективно редактировать текст. Ниже представлены основные команды, связанные с выходом из редактора, а также горячие клавиши для выполнения этих действий.

## Основные команды выхода

- Выход с сохранением изменений
- - Команда: `:wq`
- - Горячая клавиша: `ZZ` (`Shift + z + z`)
- - Описание: Сохраняет изменения и закрывает редактор.

- Выход без сохранения изменений
- - Команда: `:q!`
- - Горячая клавиша: `ZQ` (`Shift + z + q`)
- - Описание: Закрывает редактор без сохранения изменений.
- Сохранение текущего файла
- - Команда: `:w`
- - Описание: Сохраняет изменения в текущем файле.
- Сохранение файла под новым именем
- - Команда: `:w` `filename`
- - Описание: Сохраняет текущий файл под указанным именем.

## Как использовать команды

- Для выполнения команд необходимо перейти в командный режим:
- - Нажмите клавишу `Esc`, чтобы выйти из режима редактирования.
- - Затем введите команду, начиная с двоеточия `:` (например, `:wq`), и нажмите `Enter`.
- Примечания
- - Если вы хотите выйти из редактора, но не хотите сохранять изменения, используйте команду `:q!`. Эта команда игнорирует все несохраненные изменения и закрывает редактор.
- - Команды могут быть сокращены; например, вместо полного написания команды можно использовать только первые буквы (например, `:w` вместо `:write`).
- Дополнительные полезные команды
- - Отмена последнего действия: `u`
- - Повтор последнего отмененного действия: `Ctrl + r`
- - Удаление строки: `dd`
- - Копирование строки: `yy`
- - Вставка строки из буфера:
- - - Под курсором: `p`
- - - Над курсором: `P`
