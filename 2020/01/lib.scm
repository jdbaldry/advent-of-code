;; Load input from file to produce a list of numbers.
(define input
  (with-input-from-file "input.txt"
  (lambda ()
    (let loop ((lines '())
               (next-line (read-line)))
      (if (eof-object? next-line)
          (reverse lines)
          (loop (cons (string->number next-line) lines)
                (read-line)))))))

;; Print an object followed by a new line.
(define (println a) (display a) (newline))

;; Predicate that checks if the sum of the elements of a list or pair is equal to 2020.
(define (sums-to-2020? lst)
  (cond ((null? lst) #f)
        ((list? lst) (equal? (apply + lst) 2020))
        ((pair? lst) (equal? (+ (car lst) (cdr lst)) 2020))
        (else #f)))

;; Compute the cartesian product of two lists.
(define (cartesian-product a b)
  (apply append
         (map (lambda (x)
                (map (lambda (y)
                       (cons x y))
                     b))
              a)))

(define (curried-cartesian-product a)
  (lambda (b)
    (cartesian-product a b)))
