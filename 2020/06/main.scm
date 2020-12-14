(load "lib.scm")

(display "part one: ")
(display
 (fold + 0 (map (lambda (g)
                  (length
                   (delete-duplicates
                    (fold (lambda (p acc)
                            (append (string->list p) acc))
                          '()
                          g))))
                input)))
(newline)
(display "part two: ")
(display
 (fold + 0 (map (lambda (g)
                  (length
                   (filter (lambda (ayes) (equal? (length ayes) (length g)))
                           (group
                            (sort
                             (fold (lambda (p acc)
                                     (append (string->list p) acc))
                                   '()
                                   g)
                             (lambda (x y) (char<? x y)))))))
                  input)))
(newline)
