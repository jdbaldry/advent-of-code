(define input
  (with-input-from-file "input.txt"
    (lambda ()
      (let loop ((lines '())
                 (next-line (read-line)))
        (if (eof-object? next-line)
            lines
            (loop (cons next-line lines) (read-line)))))))

(define-structure bag adj col contents)

;; Split a string on a substring into a list of strings.
(define (split-string str substr)
  (let loop ((lst (string->list str))
             (sublst (string->list substr))
             (strings '())
             (current '())
             (lookahead '()))
    (if (null? lst)
        (reverse (cons (list->string (reverse current)) strings))
        (if (null? sublst)
            (loop lst
                  (string->list substr)
                  (cons (list->string (reverse current)) strings)
                  '()
                  '())
            (if (char=? (car lst) (car sublst))
                (loop (cdr lst)
                      (cdr sublst)
                      strings
                      current
                      (cons (car sublst) lookahead))
                (loop (cdr lst)
                      (string->list substr)
                      strings
                      (append (list (car lst)) lookahead current)
                      '()))))))

(define (string->name str)
  (define (words str) (split-string str " "))
  (string-append (car (words str)) " " (cadr (words str))))

;; Take a rule from input and convert it into a bag.
;; It incorrectly parses "contains no other bags" to a pair of (#f . "other bags") but that's mostly fine.
(define (string->bag str)
  (define (words str) (split-string str " "))
  (make-bag (car (words (car (split-string str "contain "))))
            (cadr (words (car (split-string str "contain "))))
            (map (lambda (str) (cons (string->number (car (words str)))
                                (string-append (cadr (words str)) " " (caddr (words str)))))
                 (split-string (cadr (split-string str "contain ")) ", "))))

(define (name bag)
  (string-append (bag-adj bag) " " (bag-col bag)))

(define (pretty-print bag)
  (display (string-append (bag-adj bag) " " (bag-col bag) ": "))(display (bag-contents bag))(newline))

;; Return the list of bags that directly contain bag-name.
(define (directly-contains bags bag-name)
  (filter (lambda (bag) (any (lambda (bag-pair) (equal? bag-name (cdr bag-pair)))
                         (bag-contents bag)))
          bags))

;; Return the list of bags that directly contain an bag-name in bag-name-list.
(define (directly-contains-any bags bag-name-list)
  (filter (lambda (bag) (any (lambda (bag-pair) (any (lambda (bag-name) (equal? bag-name (cdr bag-pair))) bag-name-list))
                        (bag-contents bag)))
          bags))

(define graph (map (lambda (line) (string->bag line)) input))

(define (graph->dot bag-list)
  (display "digraph Bags {")(newline)
  (for-each (lambda (bag)
              (for-each (lambda (content)
                          (display (string-append "    \"" (name bag) "\" -> \"" (cdr content) "\";"))(newline))
                        (bag-contents bag)))
            bag-list)
  (display "}")(newline))

;; find-bags finds bags that are able to directly or indirectly contain bag-name.
;; There must be a better way...
(define (find-bags bags bag-name)
  (let loop ((known (directly-contains bags bag-name))
             (prev '())
             (iterations 0))
    (if (equal? (length known) (length prev))
        known
        (if (< iterations 100)
            (loop (delete-duplicates (append (directly-contains-any bags (map name known)) known))
                  known
                  (+ 1 iterations))
            known))))

(define bag-table
  (let ((table (make-eq-hash-table (length input))))
    (for-each (lambda (line)
                (let ((bag (string->bag line)))
                  (hash-table/put! table (string->symbol (name bag)) bag)))
              input)
    table))

;; Return the total number of bags that must be contained within the bag.
(define (total-contents bag-name)
  (let loop ((to-check (bag-contents (hash-table/get bag-table (string->symbol bag-name) #f)))
             (total 0))
    (if (null? to-check)
        total
        ;; Turns out to be handy that we incorrectly parse "contains no bags" as (#f . "other bags")
        (if (caar to-check)
            (loop (cdr to-check)
                  (+ total (caar to-check) (* (caar to-check) (total-contents (cdar to-check)))))
            total))))
