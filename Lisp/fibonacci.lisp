(defun fibonacci (n)
  (if (or (= n 0) (= n 1))
      n
      (let ((a (fibonacci (- n 1)))
            (b (fibonacci (- n 2))))
        (+ a b))))
