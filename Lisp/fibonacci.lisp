;;;; The Fibonacci sequence, by every algorithm.
;;;;
;;;; Every function returns the same value for the same input, using the
;;;; standard indexing: F(0) = 0, F(1) = 1. Common Lisp integers are bignums,
;;;; so every implementation except the closed form is exact for any n.

;;; Transcribe the definition directly. O(phi^n): each call re-derives both
;;; subtrees, recomputing the same values exponentially often.
(defun fib-recursive (n)
  (if (or (= n 0) (= n 1))
      n
      (+ (fib-recursive (- n 1))
         (fib-recursive (- n 2)))))

;;; The recursive version with each F(n) remembered the first time it is
;;; computed, which collapses the call tree to O(n).
(defun fib-memoized (n)
  (let ((cache (make-hash-table)))
    (setf (gethash 0 cache) 0
          (gethash 1 cache) 1)
    (labels ((go-fib (k)
               (or (gethash k cache)
                   (setf (gethash k cache)
                         (+ (go-fib (- k 1)) (go-fib (- k 2)))))))
      (go-fib n))))

;;; Iterate upward keeping only the last two values. O(n) time, O(1) space.
(defun fib-linear (n)
  (let ((a 0) (b 1))
    (dotimes (i n a)
      (psetf a b
             b (+ a b)))))

;;; Multiply two 2x2 matrices, each passed as (list a b c d) in row-major order.
(defun mat-mul (m o)
  (destructuring-bind (a b c d) m
    (destructuring-bind (e f g h) o
      (list (+ (* a e) (* b g))
            (+ (* a f) (* b h))
            (+ (* c e) (* d g))
            (+ (* c f) (* d h))))))

;;; Raise [[1,1],[1,0]] to the n'th power; F(n) is the top-right entry.
;;; O(log n) via exponentiation by squaring.
(defun fib-matrix (n)
  (let ((result (list 1 0 0 1))
        (base (list 1 1 1 0))
        (e n))
    (loop while (> e 0) do
      (when (oddp e)
        (setf result (mat-mul result base)))
      (setf e (ash e -1))
      (when (> e 0)
        (setf base (mat-mul base base))))
    (second result)))

;;; Return (values F(n) F(n+1)), halving n on each step, using
;;;   F(2k)   = F(k) * (2*F(k+1) - F(k))
;;;   F(2k+1) = F(k)^2 + F(k+1)^2
;;; which is the matrix method with the algebra worked out by hand. O(log n).
(defun fib-pair (n)
  (if (= n 0)
      (values 0 1)
      (multiple-value-bind (a b) (fib-pair (ash n -1))
        (let ((c (* a (- (* 2 b) a)))
              (d (+ (* a a) (* b b))))
          (if (oddp n)
              (values d (+ c d))
              (values c d))))))

(defun fib-fast-doubling (n)
  (values (fib-pair n)))

;;; Binet's formula: F(n) = (phi^n - psi^n) / sqrt(5). The only implementation
;;; that is not exact: it rounds a double-float, and the error compounds with n
;;; until the rounding lands on the wrong integer. Uses double-floats (5d0) for
;;; the widest precision Common Lisp guarantees.
(defun fib-closed-form (n)
  (let* ((sqrt5 (sqrt 5d0))
         (phi (/ (+ 1d0 sqrt5) 2d0))
         (psi (/ (- 1d0 sqrt5) 2d0)))
    (round (/ (- (expt phi n) (expt psi n)) sqrt5))))
