(load (merge-pathnames "fibonacci.lisp" *load-truename*))

;; (n . F(n)) with the standard indexing: F(0) = 0, F(1) = 1.
(defparameter *cases*
  '((0 . 0) (1 . 1) (2 . 1) (3 . 2) (4 . 3) (5 . 5) (6 . 8)
    (7 . 13) (8 . 21) (9 . 34) (10 . 55) (20 . 6765) (30 . 832040)))

(let ((failed 0))
  (dolist (case *cases*)
    (let* ((n (car case))
           (want (cdr case))
           (got (fibonacci n)))
      (unless (= got want)
        (incf failed)
        (format t "FAIL: (fibonacci ~d) => ~d, want ~d~%" n got want))))
  (if (zerop failed)
      (format t "ok: ~d cases passed~%" (length *cases*))
      (progn
        (format t "~d of ~d cases failed~%" failed (length *cases*))
        (sb-ext:exit :code 1))))
