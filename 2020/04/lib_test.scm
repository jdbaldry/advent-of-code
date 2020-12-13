(load "lib.scm")

(define (print o) (display o) (display " ") o)
(define (test-split-string)
  (display "split-string: ")
  (print (equal? (split-string "abc  def" "  ") '("abc" "def")))
  (print (equal? (split-string "uv\n\nwx\n\nyz" "\n\n") '("uv" "wx" "yz")))
  (newline))

(define (test-read-passports)
  (display "read-passports: ")
  (let ((input "byr:2010 pid:#1bb4d8 eyr:2021 hgt:186cm iyr:2020 ecl:grt

pid:937877382 eyr:2029
ecl:amb hgt:187cm iyr:2019
byr:1933 hcl:#888785")
        (want '("byr:2010 pid:#1bb4d8 eyr:2021 hgt:186cm iyr:2020 ecl:grt" "pid:937877382 eyr:2029 ecl:amb hgt:187cm iyr:2019 byr:1933 hcl:#888785")))
    (let ((got (with-input-from-string input read-passports)))
      (print (and
              (equal? (length got) 2)
              (equal? got want)))))
  (newline))

(define (test-has-all-required-fields?)
  (display "has-all-required-fields?: ")
  (let ((input "byr:2010 pid:#1bb4d8 eyr:2021 hgt:186cm iyr:2020 ecl:grt

pid:937877382 eyr:2029
ecl:amb hgt:187cm iyr:2019
byr:1933 hcl:#888785

ecl:hzl
eyr:2020
hcl:#18171d
iyr:2019 hgt:183cm
byr:1935

hcl:#7d3b0c hgt:183cm cid:135
byr:1992 eyr:2024 iyr:2013 pid:138000309
ecl:oth

ecl:hzl
hgt:176cm pid:346059944 byr:1929 cid:150 eyr:1924 hcl:#fffffd iyr:2016"))
    (let ((got (filter (lambda (p) (has-all-required-fields? p))
                       (map string->passport (with-input-from-string input read-passports)))))
      (print (equal? (length got) 3))))
  (newline))

(define (test-valid-hgt?)
  (display "valid-hgt?: ")
  (print (equal? (valid-hgt? "72in") #t))
  (print (equal? (valid-hgt? "187cm") #t))
  (print (equal? (valid-hgt? "103") #f))
  (newline))

(test-split-string)
(test-read-passports)
(test-has-all-required-fields?)
(test-valid-hgt?)
