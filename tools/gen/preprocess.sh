cat D2DBaseTypes.h D2D1.h DXGIFormat.h | \
	grep -v "^#include.*$" | \
	gcc -E -xc -D__cplusplus - | \
	grep -v "^\s*$" | \
	grep -v "^#.*$" > \
	D2D1_preprocessed.txt

