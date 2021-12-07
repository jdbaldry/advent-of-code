(load "lib.scm")

(define (print o) (display o) (display " ") o)

(define (test-infix-eval)
  (let ((env (make-top-level-environment)))
    (display "infix-eval: ")
    (print (equal? (infix-eval '(1 + 2) env) 3))
    (print (equal? (infix-eval '((1 + 2) * (3 + 4)) env) 21))
    ;; infix-eval ignores operator precedence.
    (print (equal? (infix-eval '(1 + 2 * 3) env) 9))
    (newline)))

(define (test-lex-char)
  (display "lex-char: ")
  (let ((want (make-lexd 1 (string->list "a") (string->list " b c")))
        (got (lex-char (string->list "a b c") #\a)))
    (if (equal-fields? got want)
        (display #t)
        (begin (display #f)
               (newline)
               (display " got: ")(display-record got)
               (display "want: ")(display-record want))))
  (newline))

(test-infix-eval)
(test-lex-char)
