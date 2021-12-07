;; Take a list and make it a cyclical list.
(define (list->c-list lst)
  (set-cdr! (last-pair lst) (delay lst))
  lst)

;; Read input from file and organise as a list of lists of characters.
(define input
  (with-input-from-file "input.txt"
    (lambda ()
      (let loop ((lines '())
                 (next-line (read-line)))
        (if (eof-object? next-line)
            (reverse lines)
            (loop (cons (list->c-list (string->list next-line)) lines)
                  (read-line)))))))

(define (println o) (display o) (newline))

(define (print-l lst)
  (do ((lst lst (cdr lst)))
      ((null? c-list) (newline))
    (display (car lst))))

;; Print n elements of a cyclical list.
(define (print-n n c-list)
  (display "(")
  (do ((i 0 (+ i 1))
       (c-list c-list (if (promise? c-list) (cdr (force c-list)) (cdr c-list))))
      ((or (>= i n) (null? c-list)) (display ")") (newline))
    (display
     (if (promise? c-list)
         (car (force c-list))
         (car c-list)))
    (display " ")))

;; Define a step structure that defines how to step through a two dimensional list.
(define (make-step x y) (cons x y))
(define (step-x step) (car step))
(define (step-y step) (cdr step))

(define (is-tree? c) (equal? c #\#))

;; Return the nth element of a list that may have delayed cdrs.
(define (list-n n lst)
  (do ((i 0 (+ i 1))
       (lst lst (if (promise? lst) (cdr (force lst)) (cdr lst))))
      ((or (>= i n) (null? lst))
       (cond ((null? lst) '())
             ((promise? lst) (car (force lst)))
             (else (car lst))))))

;; Traverse a list by step.
(define (traverse lst step)
  (do ((y 0 (+ (step-y step) y))
       (x 0 (+ (step-x step) x)) (trees 0))
      ((>= y (length lst)) trees)
    (if (is-tree? (list-n x (list-n y lst)))
        (set! trees (+ trees 1)))))

;; Steps used in part two.
(define steps (list
               (make-step 1 1)
               (make-step 3 1)
               (make-step 5 1)
               (make-step 7 1)
               (make-step 1 2)))
