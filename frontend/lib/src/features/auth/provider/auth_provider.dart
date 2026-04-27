import 'package:dio/dio.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../../api/api_client.dart';
import 'auth_state.dart';

// 最新版推荐直接继承 Notifier (如果不需要代码生成的话)
class AuthNotifier extends Notifier<AuthState> {
  
  // 在 Notifier 中，初始化状态写在 build 方法里
  @override
  AuthState build() {
    return AuthState.initial();
  }

  Future<void> login(String username, String password) async {
    state = AuthState(status: AuthStatus.loading);
    try {
      // 在新版中，可以直接通过 ref 拿到其他 provider，不需要构造函数注入
      final dio = ref.read(dioProvider);
      
      final response = await dio.post('/auth/login', data: {
        'username': username,
        'password': password,
      });

      final token = response.data['token'];
      await ref.read(storageProvider).write(key: 'jwt_token', value: token);

      state = AuthState(status: AuthStatus.authenticated);
    } catch (e) {
      state = AuthState(status: AuthStatus.error, errorMessage: e.toString());
    }
  }
}

// 对应的 Provider 也要换成 NotifierProvider
final authProvider = NotifierProvider<AuthNotifier, AuthState>(() {
  return AuthNotifier();
});