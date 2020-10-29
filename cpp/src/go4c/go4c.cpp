#define GO_C_LIB_EXPORTS

#include "go4c.h"
#include <stdio.h>
#include <iostream>
#include <string>

#if 0

#include "msg.pb.h"

#pragma comment(lib, "third-party/protobuf/lib/libprotobuf.lib")
#pragma comment(lib, "third-party/protobuf/lib/libprotobuf-lite.lib")
#pragma comment(lib, "third-party/protobuf/lib/libprotoc.lib")

#endif // _MSC_VER

GO_C_LIB_API int32_t Go4CInit_C(FnCallback_C callback)
{
    printf("hello world\n");
    callback("hello", 4);

    return 0;
}

GO_C_LIB_API int32_t Go4CRelease_C()
{
    return 0;
}

GO_C_LIB_API int32_t Go4CInitCommand_C(char* data, int32_t len)
{
#if 0
    go4c_proto::CmdLogin login;
    if (login.ParseFromArray(data, len)) {

        std::string s = login.DebugString();
        printf("%s\n", s.data());
    }
    else {
        printf("parse fail\n");
    }
#else
    printf("hello cmd: %s\n", data);
#endif
    return 0;
}
