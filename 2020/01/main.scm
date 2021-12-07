(load "lib.scm")

(display "part one: ")
;; Compute the product of two numbers that add to 2020 by filtering the cartesian product of the input twice.
(println
 (first (map (lambda (pair) (* (car pair) (cdr pair)))
      (filter sums-to-2020? (cartesian-product input input)))))

(display "part two: ")
;; Compute the product of three numbers that add to 2020 using brute force.
(let ((results '()))
  (for-each (lambda (x)
              (for-each (lambda (y)
                          (for-each (lambda (z)
                                      (let ((fnd (list x y z)))
                                        (if (equal? (fold + 0 fnd) 2020)
                                            (set! results (cons (fold * 1 fnd) results)))))
                                    input))
                        input))
            input)
(println (first results)))
