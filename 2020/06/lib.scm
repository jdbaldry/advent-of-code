(define input
  (with-input-from-file "input.txt"
    (lambda ()
      (let loop ((groups '()) (next-line (read-line)) (group '()))
        (if (eof-object? next-line)
            (cons group groups)
            (if (zero? (string-length next-line))
                (loop (cons group groups) (read-line) '())
                (loop groups (read-line) (cons next-line group))))))))

;; Return a list of lists such that the concatenation of the result is equal to the list argument. Moreover, each sublist in the result contains only equal elements.
(define (group lst)
  (let loop ((groups '())
             (group '())
             (lst lst))
    (if (null? lst)
        (reverse (cons group groups))
        (if (null? group)
            (loop groups (list (car lst)) (cdr lst))
            (if (equal? (car lst) (car group))
                (loop groups (cons (car lst) group) (cdr lst))
                (loop (cons group groups) (list (car lst)) (cdr lst)))))))
