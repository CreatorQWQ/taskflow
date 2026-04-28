import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../provider/auth_provider.dart';
import '../provider/auth_state.dart';

class LoginScreen extends ConsumerWidget {
  LoginScreen({super.key});

  // 使用 final 并在 build 外定义，防止重复创建
  final _userController = TextEditingController();
  final _passController = TextEditingController();

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final authState = ref.watch(authProvider);

    return Scaffold(
      appBar: AppBar(title: const Text("登录 TaskFlow")),
      body: Padding(
        padding: const EdgeInsets.all(24.0),
        child: Center(
          child: SingleChildScrollView( // 防止键盘弹出遮挡
            child: Column(
              children: [
                TextField(
                  controller: _userController,
                  decoration: const InputDecoration(
                    labelText: "用户名",
                    prefixIcon: Icon(Icons.person),
                    border: OutlineInputBorder(),
                  ),
                ),
                const SizedBox(height: 16),
                TextField(
                  controller: _passController,
                  decoration: const InputDecoration(
                    labelText: "密码",
                    prefixIcon: Icon(Icons.lock),
                    border: OutlineInputBorder(),
                  ),
                  obscureText: true,
                ),
                const SizedBox(height: 24),
                if (authState.status == AuthStatus.loading)
                  const CircularProgressIndicator()
                else
                  SizedBox(
                    width: double.infinity,
                    height: 50,
                    child: ElevatedButton(
                      onPressed: () {
                        ref.read(authProvider.notifier).login(
                          _userController.text,
                          _passController.text,
                        );
                      },
                      child: const Text("立即登录"),
                    ),
                  ),
                if (authState.status == AuthStatus.error)
                  Padding(
                    padding: const EdgeInsets.only(top: 16),
                    child: Text(
                      authState.errorMessage ?? "登录失败",
                      style: const TextStyle(color: Colors.red),
                    ),
                  ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}