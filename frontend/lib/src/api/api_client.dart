import 'package:dio/dio.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

// 1. 创建存储对象 (加密存储 Token)
final storageProvider = Provider((ref) => const FlutterSecureStorage());

// 2. 创建 Dio 实例 Provider
final dioProvider = Provider((ref) {
  final storage = ref.watch(storageProvider);

  // 初始化 Dio 配置
  final dio = Dio(
    BaseOptions(
      // 重要提示：
      // Android 模拟器访问电脑后端请用 10.0.2.2
      // iOS 模拟器请用 127.0.0.1
      // 真机调试请用你电脑的局域网 IP (例如 192.168.1.5)
      baseUrl: 'http://192.168.0.3:8080/api/v1',
      connectTimeout: const Duration(seconds: 10), // 10 秒连接超时
      receiveTimeout: const Duration(seconds: 5), // 5 秒接收超时
    ),
  );

  // 3. 添加拦截器 (Interceptors)
  // 这是生产级应用的精髓：在请求发出前自动“加料”
  dio.interceptors.add(
    InterceptorsWrapper(
      onRequest: (options, handler) async {
        // 从“加密保险箱”里读取 JWT Token
        final token = await storage.read(key: 'jwt_token');
        
        // 如果 Token 存在，就塞进 Header
        if (token != null) {
          options.headers['Authorization'] = 'Bearer $token';
        }
        
        print("发送请求: ${options.path}");
        return handler.next(options); // 继续执行请求
      },
      onError: (DioException e, handler) {
        // 这里可以统一处理错误，比如 401 说明登录过期了
        if (e.response?.statusCode == 401) {
          print("登录已过期，需要重新登录");
        }
        return handler.next(e);
      },
    ),
  );

  return dio;
});