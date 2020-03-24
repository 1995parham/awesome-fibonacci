(defun fibonacci (n)
    (if (or (= n 0) (= n 1))
        1
        
        (let
            (
             (first (fibonacci (- n 1)))
             (second (fibonacci (- n 2))))
             (+ first second)))
            )
