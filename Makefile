challenge='01'

comp:
	clang c/2016/day_01.c -o c/2016/day01
	./c/2016/day01

c16:
	bash -c "clang ./c/2016/day_${challenge}.c -o main"
	./main
