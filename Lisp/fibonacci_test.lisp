(load (merge-pathnames "fibonacci.lisp" *load-truename*))

;; (n . F(n)) with the standard indexing: F(0) = 0, F(1) = 1.
(defparameter *cases*
  '((0 . 0) (1 . 1) (2 . 1) (3 . 2) (4 . 3) (5 . 5) (6 . 8)
    (7 . 13) (8 . 21) (9 . 34) (10 . 55) (20 . 6765) (30 . 832040)))

;; Every implementation, paired with its name for reporting.
(defparameter *implementations*
  (list (cons "recursive" #'fib-recursive)
        (cons "memoized" #'fib-memoized)
        (cons "linear" #'fib-linear)
        (cons "matrix" #'fib-matrix)
        (cons "fast-doubling" #'fib-fast-doubling)
        (cons "closed-form" #'fib-closed-form)))

;; Everything except recursive (too slow) and closed-form (not exact far out).
(defparameter *exact-implementations*
  (list (cons "memoized" #'fib-memoized)
        (cons "linear" #'fib-linear)
        (cons "matrix" #'fib-matrix)
        (cons "fast-doubling" #'fib-fast-doubling)))

(defparameter *failed* 0)

(defun check (name got want context)
  (unless (= got want)
    (incf *failed*)
    (format t "FAIL: ~a ~a => ~a, want ~a~%" name context got want)))

;; Every implementation against the shared table.
(dolist (impl *implementations*)
  (dolist (case *cases*)
    (check (car impl) (funcall (cdr impl) (car case)) (cdr case)
           (format nil "(fib ~d)" (car case)))))

;; The implementations against each other over a contiguous range, bounded by
;; recursive's exponential cost.
(dotimes (n 30)
  (let ((want (fib-linear n)))
    (dolist (impl *implementations*)
      (check (car impl) (funcall (cdr impl) n) want (format nil "(fib ~d)" n)))))

;; Bignums make the exact implementations exact arbitrarily far out. F(200) has
;; 42 digits and would overflow any fixed-width integer many times over.
(let ((want (fib-linear 200)))
  (check "linear" want 280571172992510140037611932413038677189525 "(fib 200)")
  (dolist (impl *exact-implementations*)
    (check (car impl) (funcall (cdr impl) 200) want "(fib 200)")))

;; Where the closed form stops matching the exact answer. A documented property
;; of double-float rounding rather than a bug, but a silent change to it should
;; still fail the build. SBCL diverges at F(71), the same point as Python and
;; earlier than Go and Rust (F(76)); the limit is a property of how each
;; runtime rounds expt, not of the mathematics.
(defparameter *closed-form-exact-up-to* 70)

(dotimes (n (1+ *closed-form-exact-up-to*))
  (check "closed-form" (fib-closed-form n) (fib-linear n)
         (format nil "(fib ~d) should be exact" n)))

(let ((n (1+ *closed-form-exact-up-to*)))
  (when (= (fib-closed-form n) (fib-linear n))
    (incf *failed*)
    (format t "FAIL: closed form is now exact at F(~d); the documented limit moved~%" n)))

(if (zerop *failed*)
    (format t "ok: all implementations agree~%")
    (progn
      (format t "~d checks failed~%" *failed*)
      (sb-ext:exit :code 1)))
