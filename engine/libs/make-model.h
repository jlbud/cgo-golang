#ifndef MAKE_T2S_MODEL_H_
#define MAKE_T2S_MODEL_H_

#ifdef __cplusplus
extern "C" {
#endif  // __cplusplus

#if defined _WIN64 || defined _WIN32
#define T2S_EXPORT __declspec(dllexport)
#else
#define T2S_EXPORT __attribute__((visibility("default")))
#endif

/**
 * @brief 生成用户自定义模型
 * @param[in]   in_file   自定义规则文件
 * @param[out]  out_file  自定义规则模型输出路径
 * @return int
 * @retval 0  成功
 * @retval 1  失败
 */
T2S_EXPORT int MakeT2sModel(const char* in_file, const char* out_file);

#ifdef __cplusplus
}
#endif  // __cplusplus
#endif  // MAKE_T2S_MODEL_H_
