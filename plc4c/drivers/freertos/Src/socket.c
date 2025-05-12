//
// Created by Süleyman ÇİÇEK <suleyman@cicek.pw> on 5/10/2025.
//

#include <sys/socket.h>

#ifndef _KERNEL

__BEGIN_DECLS
int	accept(int, struct sockaddr *, socklen_t *)
{
	// TODO: Implement this function
	return 0;
}

int	bind(int, const struct sockaddr *, socklen_t)
{
	// TODO: Implement this function
	return 0;
}

int	connect(int, const struct sockaddr *, socklen_t)int	accept(int, struct sockaddr *, socklen_t *)
{
	// TODO: Implement this function
	return 0;
}

int	getpeername(int, struct sockaddr *, socklen_t *)
{
	// TODO: Implement this function
	return 0;
}

int	getsockname(int, struct sockaddr *, socklen_t *)
{
	// TODO: Implement this function
	return 0;
}

int	getsockopt(int, int, int, void *, socklen_t *)
{
	// TODO: Implement this function
	return 0;
}

int	listen(int, int)
{
	// TODO: Implement this function
	return 0;
}

ssize_t	recv(int, void *, size_t, int)
{
	// TODO: Implement this function
	return 0;
}

ssize_t	recvfrom(int, void *, size_t, int, struct sockaddr *, socklen_t *)
{
	// TODO: Implement this function
	return 0;
}

ssize_t	recvmsg(int, struct msghdr *, int)
{
	// TODO: Implement this function
	return 0;
}

int	recvmmsg(int, struct mmsghdr *, unsigned int, int, struct timespec *)
{
	// TODO: Implement this function
	return 0;
}

ssize_t	send(int, const void *, size_t, int)
{
	// TODO: Implement this function
	return 0;
}

ssize_t	sendto(int, const void *,
		      size_t, int, const struct sockaddr *, socklen_t)
{
	// TODO: Implement this function
	return 0;
}

ssize_t	sendmsg(int, const struct msghdr *, int)
{
	// TODO: Implement this function
	return 0;
}

int	sendmmsg(int, struct mmsghdr *, unsigned int, int)
{
	// TODO: Implement this function
	return 0;
}

int	setsockopt(int, int, int, const void *, socklen_t)
{
	// TODO: Implement this function
	return 0;
}

int	shutdown(int, int)
{
	// TODO: Implement this function
	return 0;
}

int	sockatmark(int)
{
	// TODO: Implement this function
	return 0;
}

int	socket(int, int, int)
{
	// TODO: Implement this function
	return 0;
}

int	socketpair(int, int, int, int *)
{
	// TODO: Implement this function
	return 0;
}


#if __BSD_VISIBLE
int	accept4(int, struct sockaddr *__restrict, socklen_t *__restrict, int);
#endif

#if __BSD_VISIBLE
int	getpeereid(int, uid_t *, gid_t *);
int	getrtable(void);
int	setrtable(int);
#endif /* __BSD_VISIBLE */

__END_DECLS

#else

static inline struct sockaddr *
sstosa(struct sockaddr_storage *ss)
{
	return ((struct sockaddr *)(ss));
}

#endif /* !_KERNEL */