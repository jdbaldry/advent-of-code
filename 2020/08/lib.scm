(define input
  (with-input-from-file "input.txt"
    (lambda ()
      (let loop ((lines '())
                 (next-line (read-line)))
        (if (eof-object? next-line)
            (reverse lines)
            (loop (cons next-line lines) (read-line)))))))

(define (string->instruction s)
  (define tokens (let loop  ((l (string->list s))
                             (tokens '())
                             (token '()))
                   (if (null? l)
                       (reverse (cons (reverse token) tokens))
                       (if (char=? #\space (car l))
                           (loop (cdr l) (cons (reverse token) tokens) '())
                           (loop (cdr l) tokens (cons (car l) token))))))
  (cons (string->symbol (list->string (car tokens))) (string->number (list->string (cadr tokens)))))

(define instructions (map string->instruction input))

;; make-machine creates a machine that will execute instructions until the step-limit is reached.
;; Instructions are expected to be a list of operation, argument pairs where an operation is a symbol and the argument is an integer number.
(define (make-machine instructions step-limit)
  (let ((acc 0)
        (pc 0)
        (steps 0)
        (limit step-limit))
    (define (step)
      (cond ((or (>= pc (length instructions)) (< pc 0)) 'halted)
            ((>= steps limit) 'stopped)
            (else (let ((op (car (list-ref instructions pc)))
                        (arg (cdr (list-ref instructions pc))))
                    (cond ((eq? op 'acc) (begin (set! acc (+ arg acc))
                                                (set! pc (+ 1 pc))))
                          ((eq? op 'jmp) (set! pc (+ arg pc)))
                          ((eq? op 'nop) (set! pc (+ 1 pc)))
                          (else (error "Unknown op -- MACHINE" op)))
                    (set! steps (+ 1 steps))
                    'ok))))

    (define (exec)
      (do ((state 'ok (step)))
          ((not (eq? state 'ok)) state)))
    (define (dispatch m)
      (cond ((eq? m 'exec) (exec))
            ((eq? m 'step) (step))
            ((eq? m 'get-acc) acc)
            ((eq? m 'get-pc) pc)
            (else (error "Unknown operation -- MACHINE" m))))
    dispatch))

;; probe-for-loop steps through a machines execution and checks whether the program counter reaches a previous state. A pair of state and acc is returned once either the execution stops or a loop is encountered.
(define (probe-for-loop step-limit)
  (let ((m (make-machine instructions step-limit)))
    ;; seen is a vector one longer than the length of instructions to check if we finish the full set.
    (do ((seen (make-vector (+ 1 (length instructions)) 0))
         (state (m 'step) (m 'step))
         (acc (m 'get-acc) (m 'get-acc))
         (pc (m 'get-pc) (m 'get-pc)))
        ((or (eq? (vector-ref seen pc) 1) (not (eq? 'ok state)))
         (if (eq? (vector-ref seen pc) 1)
             (cons 'looping acc)
             (cons state acc)))
      (vector-set! seen pc 1))))

;; flip-op flips a jmp to a nop, or a nop to a jump.
(define (flip-op op)
  (cond ((eq? op 'jmp) 'nop)
        ((eq? op 'nop) 'jmp)
        (else op)))

;; flip-op-at-ref returns a new set of instructions with the operation at ref flipped.
(define (flip-op-at-ref instructions ref)
  (map (lambda (zip) (if (equal? ref (car zip))
                    (cons (flip-op (cadr zip)) (cddr zip))
                    (cons (cadr zip) (cddr zip))))
       (zip (iota (length instructions)) instructions)))

;; Produce a list of program instructions that may resolve a looping program.
(define (possibly-correct instructions)
  (map (lambda (i) (flip-op-at-ref instructions i))
       (iota (length instructions))))

;; Try to find which possibly-correct instruction set halts the machine and returns the machines acc.
(define (find-correct instructions)
  (car (filter integer?
               (map (lambda (instructions)
                      (let ((m (make-machine instructions 500)))
                        (if (eq? (m 'exec) 'halted)
                            (m 'get-acc)
                            #f)))
                    (possibly-correct instructions)))))

;; zip-with zips two lists together using f.
;; The resulting list is the length of the shortest input list.
(define (zip-with f l1 l2)
  (if (or (null? l1) (null? l2))
      '()
      (cons (f (car l1) (car l2)) (zip-with f (cdr l1) (cdr l2)))))

;; zip zips two lists together to for a list of pairs.
;; The resulting list is the lenght of the shortest input list.
(define (zip l1 l2) (zip-with cons l1 l2))

;; vector-zip zips two vectors together.
(define (vector-zip-with) ())
