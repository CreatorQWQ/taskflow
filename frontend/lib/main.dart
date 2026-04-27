// lib/main.dart

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'src/features/auth/provider/auth_provider.dart';
import 'src/features/auth/provider/auth_state.dart';

void main() {
  runApp(const ProviderScope(child: MyApp()));
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'TaskFlow',
      home: LoginScreen(),
    );
  }
}

class LoginScreen extends ConsumerWidget {
  final _userController = TextEditingController();
  final _passController = TextEditingController();

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    // 监听 authProvider 的状态
    final authState = ref.watch(authProvider);

    return Scaffold(
      appBar: AppBar(title: const Text("CreatorQWQ, 登录 TaskFlow")),
      body: Padding(
        padding: const EdgeInsets.all(20.0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            TextField(controller: _userController, decoration: const InputDecoration(labelText: "用户名")),
            TextField(controller: _passController, decoration: const InputDecoration(labelText: "密码"), obscureText: true),
            const SizedBox(height: 30),
            
            // 根据不同状态显示不同内容
            if (authState.status == AuthStatus.loading)
              const CircularProgressIndicator()
            else
              ElevatedButton(
                onPressed: () {
                  // 调用 Notifier 里的登录方法
                  ref.read(authProvider.notifier).login(
                    _userController.text,
                    _passController.text,
                  );
                },
                child: const Text("登录"),
              ),

            // 如果报错，显示错误信息
            if (authState.status == AuthStatus.error)
              Padding(
                padding: const EdgeInsets.only(top: 20),
                child: Text(authState.errorMessage ?? "错误", style: const TextStyle(color: Colors.red)),
              ),

            // 如果成功，显示成功提示
            if (authState.status == AuthStatus.authenticated)
              const Padding(
                padding: const EdgeInsets.only(top: 20),
                child: Text("登录成功！", style: TextStyle(color: Colors.green, fontSize: 20)),
              ),
          ],
        ),
      ),
    );
  }
}