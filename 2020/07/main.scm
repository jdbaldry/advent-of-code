(load "lib.scm")

(display "part one: ")
(display (length (find-bags (map string->bag input) "shiny gold")))
(newline)
(display "part two: ")
(display (total-contents "shiny gold"))
(newline)
