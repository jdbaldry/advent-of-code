;; Load a list of lines from the input file.
(define input
  (with-input-from-file "input.txt"
    (lambda ()
      (let loop ((lines '())
                 (next-line (read-line)))
        (if (eof-object? next-line)
            (reverse lines)
            (loop (cons next-line lines)
                  (read-line)))))))

(define (seat-id row column)
  (+ column (* 8 row)))

(define (string->bit-string str)
  (do ((str (reverse (string->list str)) (cdr str))
       (i 0 (+ i 1))
       (bstr (make-bit-string (string-length str) #f)))
      ((null? str) bstr)
    (if (or (char=? (car str) #\B) (char=? (car str) #\R))
        (bit-string-set! bstr i)
        (bit-string-clear! bstr i))))

(define (row bit-string)
  (bit-string->unsigned-integer (bit-substring bit-string 3 10)))

(define (column bit-string)
  (bit-string->unsigned-integer (bit-substring bit-string 0 3)))

;; Find a missing number in a sorted list of consecutive numbers.
(define (missing-number lst)
  (do ((lst lst (cdr lst))
       (prev '() (car lst)))
      ((or (null? lst)
           (if (null? prev)
               #f
               (not (equal? prev (+ 1 (car lst)))))) (- prev 1))))
