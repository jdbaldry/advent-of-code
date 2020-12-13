(load "lib.scm")

;; Print an object and return it.
(define (print o) (display o) (display " ") o)

(define (test-split-string)
  (display "split-string: ")
  (print (equal? (split-string "abc def" #\space) '("abc" "def")))
  (print (equal? (split-string "tu-vw-xyz" #\-) '("tu" "vw" "xyz")))
  (newline))

(define (test-string->entry)
  (display "string->entry: ")
  (let ((entry (string->entry "1-2 c: password")))
    (print
     (and
      (equal? (entry-min entry) 1)
      (equal? (entry-max entry) 2)
      (equal? (entry-character entry) #\c)
      (equal? (entry-password entry) "password"))))
  (newline))

(define (test-part-one-password-valid?)
  (display "part-one-password-valid?: ")
  (print (equal? (part-one-password-valid? (string->entry "1-3 a: abcde")) #t))
  (print (equal? (part-one-password-valid? (string->entry "1-3 b: cdefg")) #f))
  (newline))

(define (test-part-two-password-valid?)
  (display "part-two-password-valid?: ")
  (print (equal? (part-two-password-valid? (string->entry "1-3 a: abcde")) #t))
  (print (equal? (part-two-password-valid? (string->entry "1-3 b: cdefg")) #f))
  (print (equal? (part-two-password-valid? (string->entry "2-9 c: ccccccccc")) #f))
  (newline))

(define (test-matching-char?)
  (display "matching-char?: ")
  (print (equal? (matching-char? #\a 0 "a") #t))
  (print (equal? (matching-char? #\b 1 "ab") #t))
  (print (equal? (matching-char? #\c 2 "ccccccccc") #t))
  (print (equal? (matching-char? #\c 9 "cccccccccc") #t))
  (newline))


(test-split-string)
(test-string->entry)
(test-part-one-password-valid?)
(test-part-two-password-valid?)
(test-matching-char?)
