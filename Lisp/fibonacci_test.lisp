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

;; Report where the closed form stops matching the exact answer. This is a
;; measured property of double-float rounding; the value is printed rather than
;; asserted because it differs across languages (75 in Go and Rust, 70 in
;; Python) and is not known here until this test runs.
(let ((limit nil))
  (loop for n from 0 below 200
        while (= (fib-closed-form n) (fib-linear n))
        do (setf limit n))
  (format t "closed-form: exact through F(~d)~%" limit)
  ;; A conservative floor that any reasonable double-float implementation clears.
  (when (< limit 60)
    (incf *failed*)
    (format t "FAIL: closed form only exact to F(~d), expected at least F(60)~%" limit)))

(if (zerop *failed*)
    (format t "ok: all implementations agree~%")
    (progn
      (format t "~d checks failed~%" *failed*)
      (sb-ext:exit :code 1)))
