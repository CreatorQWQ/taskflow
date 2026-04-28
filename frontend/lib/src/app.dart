import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'features/auth/provider/auth_provider.dart';
import 'features/auth/provider/auth_state.dart';
import 'features/auth/view/login_screen.dart';
import 'features/tasks/view/home_screen.dart';

class TaskFlowApp extends ConsumerWidget {
  const TaskFlowApp({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final authState = ref.watch(authProvider);

    // 如果正在检查 Token，显示一个全屏加载
    if (authState.status == AuthStatus.checking) {
      return const MaterialApp(
        home: Scaffold(body: Center(child: CircularProgressIndicator())),
      );
    }

    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'TaskFlow',
      theme: ThemeData(useMaterial3: true, colorSchemeSeed: Colors.blue),
      home: authState.status == AuthStatus.authenticated
          ? const HomeScreen()
          : LoginScreen(),
    );
  }
}