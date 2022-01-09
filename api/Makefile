.PHONY: server clean
server:
	make -C server build
	mkdir -p build
	mv server/build/server build/

clean:
	make -C server clean
	rm -rf build
