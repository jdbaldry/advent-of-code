(load "lib.scm")

(define sorted-seats
  (sort (map (lambda (str)
               (seat-id (row (string->bit-string str))
                        (column (string->bit-string str))))
             input)
        (lambda (x y) (> x y))))
(display "part one: ")
(display (first sorted-seats))
(newline)

(display "part two: ")
(display (missing-number sorted-seats))
(newline)
