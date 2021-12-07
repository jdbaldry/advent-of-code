(load "lib.scm")

(display "part one: ")
(display (cdr (probe-for-loop 250)))
(newline)
(display "part two: ")
(display (find-correct instructions))
(newline)
