(load "lib.scm")

(display "part one: ")
(println (traverse input (make-step 3 1)))

(display "part two: ")
(println (fold * 1 (map (lambda (step) (traverse input step)) steps)))
