dist_dir := dist
cc := go build
libs := hex

all : pkt-recv

pkt-recv: pkt-recv/pkt-recv.go ${libs}
	${cc} -o ${dist_dir}/pkt-recv pkt-recv/pkt-recv.go

clean:
	rm ${dist_dir}/*
