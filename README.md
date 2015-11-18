# transliterate

Remove accents from stdin and writes the output to stdout

Example:

```bash
$ echo "Maître Corbeau, sur un arbre perché," | transliterate
Maitre Corbeau, sur un arbre perche,
```

## Credits

* Implementation: Golang blog [text normalization][art] article
* The name "transliterate": Rails' [`ActiveSupport::Inflector.transliterate`][ar]

[art]: https://blog.golang.org/normalization
[ar]: http://api.rubyonrails.org/classes/ActiveSupport/Inflector.html#method-i-transliterate
