/**
 * @file t2s.h
 * @brief text to symbol
 * @details
 * @authors Yan Qibao, Li Peng
 * @copyright 2020 Unisound AI Technology Co., Ltd. All rights reserved.
 */

#ifndef TEXT_TO_SYMBOL_T2S_H_
#define TEXT_TO_SYMBOL_T2S_H_

#if defined _WIN64 || defined _WIN32
#define T2S_EXPORT __declspec(dllexport)
#else
#define T2S_EXPORT __attribute__((visibility("default")))
#endif

#ifdef __cplusplus
extern "C" {
#endif

typedef enum {
  T2S_SUCCESS,            /**< 正确返回 */
  T2S_ERROR,              /**< 一般错误 */
  T2S_DATA_ERROR,         /**< 参数错误 */
  T2S_FILE_NOT_FOUND,     /**< 未找到文件 */
  T2S_FILE_CANNOT_OPEN,   /**< 不能打开文件 */
  T2S_MEMORY_ALLOC_FAILED /**< 分配内存失败 */
} T2S_ERRNO;

/**
 * @brief 获取版本号
 * @return const char*
 */
T2S_EXPORT const char* T2sGetVersion();

/**
 * @brief 读取模型
 * @param[in]  path  模型路径
 * @return T2S_ERROR_CODE
 * @retval T2S_SUCCESS              读取模型成功
 * @retval T2S_DATA_ERROR          参数错误
 * @retval T2S_FILE_NOT_FOUND       未找到文件
 * @retval T2S_FILE_CANNOT_OPEN     不能打开文件
 * @retval T2S_MEMORY_ALLOC_FAILED  申请内存失败
 * @retval T2S_ERROR                其他原因导致读取模型失败
 */
T2S_EXPORT T2S_ERRNO LoadT2sModel(const char* path, void** model);

/**
 * @brief 初始化实例
 * @param[in]  model     模型
 * @param[out] instance  实例
 * @return T2S_ERROR_CODE
 * @retval T2S_SUCCESS     初始化实例成功
 * @retval T2S_DATA_ERROR  参数错误
 */
T2S_EXPORT T2S_ERRNO InitializeT2sInstance(const void* model, void** instance);

/**
 * @brief 销毁实例
 * @param[in,out]  instance  实例
 * @return T2S_ERROR_CODE
 * @retval T2S_SUCCESS     销毁实例成功
 * @retval T2S_DATA_ERROR  参数错误
 */
T2S_EXPORT T2S_ERRNO TerminateT2sInstance(void** instance);

/**
 * @brief 释放模型数据
 * @param[in,out]  model  模型数据
 * @return T2S_ERROR_CODE
 * @retval T2S_SUCCESS     释放模型数据成功
 * @retval T2S_DATA_ERROR  参数错误
 */
T2S_EXPORT T2S_ERRNO UnloadT2sModel(void** model);

/**
 * @brief 读取用户自定义模型
 * @param[in]      data      模型数据
 * @param[in]      size      模型数据长度
 * @param[in,out]  instance  实例
 * @return T2S_ERROR_CODE
 * @retval T2S_SUCCESS              读取模型成功
 * @retval T2S_DATA_ERROR           参数错误
 * @retval T2S_MEMORY_ALLOC_FAILED  申请内存失败
 */
T2S_EXPORT T2S_ERRNO LoadUserRules(const char* data, int size, void* instance);

/**
 * @brief 释放用户自定义模型数据
 * @param[in,out]  instance  实例
 * @return T2S_ERROR_CODE
 * @retval T2S_SUCCESS     释放模型数据成功
 * @retval T2S_DATA_ERROR  参数错误
 */
T2S_EXPORT T2S_ERRNO UnloadUserRules(void* instance);

/**
 * @brief 后处理清空之前处理数据
 * @param[in,out]  instance      实例
 * @return T2S_ERROR_CODE
 * @retval T2S_SUCCESS              操作成功
 * @retval T2S_DATA_ERROR          参数错误
 */
T2S_EXPORT T2S_ERRNO T2sReset(void* instance);

/**
 * @brief 基于模型的分词处理
 * @param[in]  instance      实例
 * @param[in]  input      输入字符串
 * @param[in]  end_flag   回话结束标志
 * @param[out] output     分词结果
 * @return T2S_ERROR_CODE
 * @retval T2S_SUCCESS              操作成功
 * @retval T2S_DATA_ERROR          参数错误
 * @retval T2S_MEMORY_ALLOC_FAILED  申请内存失败
 * @retval T2S_ERROR                其他原因导致失败
 */
T2S_EXPORT T2S_ERRNO T2sProcess(void* instance, const char* input,
                                int end_flag, char** output);

#ifdef __cplusplus
}
#endif

#endif  // TEXT_TO_SYMBOL_T2S_H_
