// g++ - fPIC - shared go4c.cpp - o libgo4c.so

#ifdef _MSC_VER

#ifdef GO_C_LIB_EXPORTS
#define GO_C_LIB_EXPORTS __declspec(dllexport)
#else
#define GO_C_LIB_EXPORTS __declspec(dllimport)
#endif

#else

#define GO_C_LIB_EXPORTS

#endif // _MSC_VER


#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif // __cplusplus

    typedef void (*FnCallback_C)(char* data, int32_t len);

    GO_C_LIB_EXPORTS int32_t Go4CInit_C(FnCallback_C callback);

    GO_C_LIB_EXPORTS int32_t Go4CRelease_C();

    GO_C_LIB_EXPORTS int32_t Go4CInitCommand_C(char* data, int32_t len);

#ifdef __cplusplus
};
#endif // __cplusplus

