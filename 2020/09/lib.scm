(define input
  (with-input-from-file "input.txt"
    (lambda ()
      (let loop ((lines '())
                 (next-line (read-line)))
        (if (eof-object? next-line)
            (reverse lines)
            (loop (cons next-line lines) (read-line)))))))

;; lexd is a structure of the count of lexd characters, the resulting token, and the list of the remaining unlexd characters.
(define-structure lexd count tokens remaining)

;; display-record displays a records type and its fields.
(define (display-record r)
  (if (not (record? r))
      (error "not a record -- DISPLAY-RECORD" r)
      (let ((rtd (record-type-descriptor r)))
        (display (list (record-type-name rtd) ":" (map (lambda (field-name) ((record-accessor rtd field-name) r)) (record-type-field-names rtd))))
        (newline))))

;; equal-fields? checks the fields of a record for equality.
(define (equal-fields? r1 r2)
  (let ((rtd (record-type-descriptor r1)))
    (and (equal? rtd (record-type-descriptor r2))
         (every (lambda (field-name) (equal? ((record-accessor rtd field-name) r1)
                                        ((record-accessor rtd field-name) r2)))
                (record-type-field-names rtd)))))

;; lex lexs a string into a pair where the first element is a token lex tree and the second is the remaining unlexd string.
(define (lex l)
  (let loop ((count 0)
             (tokens '())
             (remaining l))
    (if (null? remaining)
        (make-lexd count (reverse tokens) remaining)
        (let ((left-paren (lex-left-paren remaining))
              (right-paren (lex-right-paren remaining))
              (spaces (lex-space remaining))
              (op (lex-op remaining))
              (int (lex-int remaining)))
          (cond ((> (lexd-count left-paren) 0) (let ((subexp (lex (lexd-remaining left-paren))))
                                                   (loop (+ count (lexd-count subexp))
                                                         (cons (lexd-tokens subexp) tokens)
                                                         (lexd-remaining subexp))))
                ((> (lexd-count right-paren) 0) (make-lexd count (reverse tokens) (lexd-remaining right-paren)))
                ((> (lexd-count spaces) 0) (loop (+ count (lexd-count spaces))
                                                   tokens
                                                   (lexd-remaining spaces)))
                ((> (lexd-count op) 0) (loop (+ count (lexd-count op))
                                               (cons (lexd-tokens op) tokens)
                                               (lexd-remaining op)))
                ((> (lexd-count int) 0) (loop (+ count (lexd-count int))
                                                (cons (lexd-tokens int) tokens)
                                                (lexd-remaining int)))
                (else (make-lexd count (reverse tokens) remaining)))))))

(define (lex-char l c)
  (if (null? l)
      (make-lexd 0 '() l)
      (if (char=? c (car l))
          (make-lexd 1 (list c) (cdr l))
          (make-lexd 0 '() l))))

(define (lex-left-paren l) (lex-char l #\())
(define (lex-right-paren l) (lex-char l #\)))
(define (lex-space l) (lex-char l #\space))

;; lex-op lexs an operator from a list of characters returing a pair where the first element is the operator and the second is the list of remaining unlexd characters.
(define (lex-op l)
  (let ((remaining l))
    (if (null? remaining)
        (make-lexd 0 '() remaining)
        (cond ((char=? #\+ (car remaining)) (make-lexd 1 + (cdr remaining)))
              ((char=? #\- (car remaining)) (make-lexd 1 - (cdr remaining)))
              ((char=? #\* (car remaining)) (make-lexd 1 * (cdr remaining)))
              ((char=? #\/ (car remaining)) (make-lexd 1 / (cdr remaining)))
              (else (make-lexd 0 '() remaining))))))

(define (lex-int l)
  (let loop ((count 0)
             (int '())
             (remaining l))
    (if (null? remaining)
        (make-lexd count (string->number (list->string (reverse int))) remaining)
        (if (char-numeric? (car remaining))
            (loop (+ 1 count) (cons (car remaining) int) (cdr remaining))
            (make-lexd count (if (null? int)
                                   '()
                                   (string->number (list->string (reverse int)))) remaining)))))

;; infix-eval evaluates a infix binary expressions with number operands from left-to-right, ignoring operator precedence.
(define (infix-eval exp env)
  (let ((l (car exp))
        (op (cadr exp))
        (r (caddr exp))
        (remaining (cdddr exp)))
    (let ((result (eval (list op
                            (if (number? l) l (infix-eval l env))
                            (if (number? r) r (infix-eval r env)))
                      env)))
      (if (null? remaining)
          result
          (infix-eval (cons result remaining) env)))))

;; TODO: p-infix-eval uses infix-eval to evaluate infix binary expressions with number operands from left-to-right, where + has precedence over *.
