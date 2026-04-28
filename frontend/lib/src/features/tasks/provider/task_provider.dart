import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../../api/api_client.dart';
import '../model/task_model.dart';

class TaskNotifier extends AsyncNotifier<List<TaskModel>> {
  // 1. 获取任务列表的核心方法
  Future<List<TaskModel>> fetchTasks() async {
    final dio = ref.read(dioProvider);
    final response = await dio.get('/tasks');

    // 将返回的 List<dynamic> 转换为 List<TaskModel>
    final List data = response.data;
    return data.map((json) => TaskModel.fromJson(json)).toList();
  }

  // 2. build 方法定义初始加载逻辑
  @override
  Future<List<TaskModel>> build() async {
    return fetchTasks();
  }

  // 3. 手动刷新方法 (下拉刷新时用)
  Future<void> refresh() async {
    state = const AsyncValue.loading(); // 先展示加载中
    state = await AsyncValue.guard(() => fetchTasks()); // 抓取并自动处理错误
  }

  // 在 TaskNotifier 类中添加以下方法
  Future<void> addTask(String title, String content) async {
    final dio = ref.read(dioProvider);

    // 1. 发送 POST 请求到后端
    await dio.post('/tasks/', data: {'title': title, 'content': content});

    // 2. 【神器】让当前 Provider 失效，Riverpod 会自动重新调用 build() 去后端抓取最新列表
    ref.invalidateSelf();
  }
}

// 4. 定义全局 Provider
final taskListProvider = AsyncNotifierProvider<TaskNotifier, List<TaskModel>>(
  () {
    return TaskNotifier();
  },
);
