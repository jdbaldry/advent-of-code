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

;; Print an object followed by a new line.
(define (println a) (display a) (newline))

;; Define an entry as a list structure and provide accessors.
(define (make-entry min max character password)
  (list min max character password))
(define (entry-min entry) (car entry))
(define (entry-max entry) (cadr entry))
(define (entry-character entry) (caddr entry))
(define (entry-password entry) (cadddr entry))

;; Splits a string on a character c to return a list of strings.
(define (split-string str c)
  (let loop ((lst (string->list str))
             (tokens '()))
    (if (pair? lst)
        (if (char=? (car lst) c)
            (cons (list->string (reverse tokens)) (loop (cdr lst) '()))
            (loop (cdr lst) (cons (car lst) tokens)))
        (if (null? tokens)
            '()
            (list (list->string (reverse tokens)))))))

(define (string->entry str)
  (define words (split-string str #\space))
  (define requirements (split-string (list-ref words 0) #\-))
  (let ((min (string->number (list-ref requirements 0)))
        (max (string->number (list-ref requirements 1)))
        (character (string-ref (list-ref words 1) 0))
        (password (list-ref words (- (length words) 1))))
    (make-entry min max character password)))

;; Is a password valid according to the policy defined in part one of the puzzle?
(define (part-one-password-valid? entry)
  (let ((count (length (filter (lambda (char)
                                 (char=? char (entry-character entry)))
                               (string->list (entry-password entry))))))
    (and (>= count (entry-min entry))
         (<= count (entry-max entry)))))

;; Does the character in string str at index ref match char?
(define (matching-char? char ref str)
  (if (and (positive? ref)
           (>= ref (string-length str)))
      #f
      (equal? char (string-ref str ref))))

;; Is a password valid according to the policy defined in part two of the puzzle?
(define (part-two-password-valid? entry)
  (define (exclusive-or a b) (not (boolean=? a b)))
 (exclusive-or
   (matching-char? (entry-character entry) (- (entry-min entry) 1) (entry-password entry))
   (matching-char? (entry-character entry) (- (entry-max entry) 1) (entry-password entry))))
