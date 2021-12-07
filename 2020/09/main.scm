(load "lib.scm")

(display "part one: ")
(display
 (let ((env (make-top-level-environment)))
   (fold +
         0
         (map
          (lambda (line)
            (infix-eval (lexd-tokens (lex (string->list line))) env))
          input))))
(newline)
(display "part two: ")
;; (display
;;  (let ((env (make-top-level-environment)))
;;    (fold +
;;          0
;;          (map
;;           (lambda (line)
;;             (p-infix-eval (lexd-tokens (lex (string->list line))) env))
;;           input))))
(newline)
