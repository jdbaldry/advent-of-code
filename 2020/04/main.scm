(load "lib.scm")

(display "part one: ")
(display (length (filter (lambda (p) (has-all-required-fields? p))
                         (map string->passport input))))
(newline)

(display "part two: ")
(display (length (filter (lambda (p) (valid? p))
                         (map string->passport input))))
(newline)
