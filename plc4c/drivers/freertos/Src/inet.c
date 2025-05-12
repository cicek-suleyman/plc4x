//
// Created by Süleyman ÇİÇEK <suleyman@cicek.pw> on 5/10/2025.
//

#include <arpa/inet.h>
#include "FreeRTOS_IP.h"

in_addr_t	 inet_addr(const char *cp)
{
	return FreeRTOS_inet_addr(cp);
}

char		*inet_ntoa(struct in_addr in)
{
	static char buf[16];
	FreeRTOS_inet_ntoa(in.s_addr, buf);
	return buf;
}

const char	*inet_ntop(int af, const void *__restrict src, char *__restrict dst,
			     socklen_t size) __attribute__ ((__bounded__(__string__,3,4)))
{
	return FreeRTOS_inet_ntop(af, src, dst, size);
}

int		 inet_pton(int af, const char *__restrict src, void *__restrict dst)
{
	return FreeRTOS_inet_pton(af, src, dst);
}


#if __BSD_VISIBLE
int		 inet_aton(const char *cp, struct in_addr *inp)
{
	inp->s_addr = inet_addr(cp);
	return 0;
}

in_addr_t	 inet_lnaof(struct in_addr in)
{
	return in.s_addr; // TODO
}

struct in_addr	 inet_makeaddr(in_addr_t net, in_addr_t host)
{
	// TODO
}

char *		 inet_neta(in_addr_t src, char *dst, size_t size)
__attribute__((__bounded__(__string__,2,3)))
{
	// TODO
}

in_addr_t	 inet_netof(struct in_addr in)
{
	return in.s_addr; // TODO
}

in_addr_t	 inet_network(const char *cp)
{
	return inet_addr(cp);
}

char		*inet_net_ntop(int af, const void *src, int bits, char *dst, size_t size)
__attribute__((__bounded__(__string__,4,5)))
{
	return FreeRTOS_inet_ntop(af, src, dst, size);
}

int		 inet_net_pton(int af, const char *src, void *dst, size_t size)
__attribute__((__bounded__(__string__,3,4)))
{
	return inet_pton(af, src, dst);
}
#endif /* __BSD_VISIBLE */