import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../../auth/provider/auth_provider.dart';
import '../provider/task_provider.dart';

class HomeScreen extends ConsumerWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final taskAsync = ref.watch(taskListProvider);

    return Scaffold(
      appBar: AppBar(
        title: const Text("我的任务清单"),
        actions: [
          IconButton(
            icon: const Icon(Icons.logout),
            onPressed: () => ref.read(authProvider.notifier).logout(),
          ),
        ],
      ),
      body: taskAsync.when(
        data: (tasks) => RefreshIndicator(
          // 当用户下拉时，触发这个方法
          onRefresh: () => ref.refresh(taskListProvider.future),
          child: tasks.isEmpty
              ? const Center(
                  // 注意：如果列表为空，ListView 必须铺满全屏，下拉刷新才有效
                  child: SingleChildScrollView(
                    physics: AlwaysScrollableScrollPhysics(),
                    child: SizedBox(
                      height: 500,
                      child: Center(child: Text("暂无任务")),
                    ),
                  ),
                )
              : ListView.builder(
                  // 确保即使列表项很少，也能下拉刷新
                  physics: const AlwaysScrollableScrollPhysics(),
                  itemCount: tasks.length,
                  itemBuilder: (context, index) {
                    final task = tasks[index];
                    return ListTile(title: Text(task.title));
                  },
                ),
        ),
        loading: () => const Center(child: CircularProgressIndicator()),
        error: (err, stack) => Center(
          child: IconButton(
            icon: const Icon(Icons.refresh),
            onPressed: () => ref.invalidate(taskListProvider), // 失败点击刷新
          ),
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () => _showAddTaskSheet(context, ref),
        child: const Icon(Icons.add),
      ),
    );
  }
}

// 弹出底部抽屉的方法
void _showAddTaskSheet(BuildContext context, WidgetRef ref) {
  final titleController = TextEditingController();
  final contentController = TextEditingController();

  showModalBottomSheet(
    context: context,
    isScrollControlled: true, // 允许随键盘升起
    builder: (context) => Padding(
      padding: EdgeInsets.only(
        bottom: MediaQuery.of(context).viewInsets.bottom, // 避开键盘
        left: 20,
        right: 20,
        top: 20,
      ),
      child: Column(
        mainAxisSize: MainAxisSize.min, // 高度自适应内容
        children: [
          const Text(
            "新建任务",
            style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
          ),
          const SizedBox(height: 16),
          TextField(
            controller: titleController,
            decoration: const InputDecoration(
              labelText: "任务标题",
              border: OutlineInputBorder(),
            ),
          ),
          const SizedBox(height: 12),
          TextField(
            controller: contentController,
            decoration: const InputDecoration(
              labelText: "详细内容",
              border: OutlineInputBorder(),
            ),
            maxLines: 3,
          ),
          const SizedBox(height: 20),
          SizedBox(
            width: double.infinity,
            height: 50,
            child: ElevatedButton(
              onPressed: () async {
                if (titleController.text.isNotEmpty) {
                  // 调用我们刚才写的 addTask
                  await ref
                      .read(taskListProvider.notifier)
                      .addTask(titleController.text, contentController.text);
                  Navigator.pop(context); // 关闭抽屉
                }
              },
              child: const Text("创建任务"),
            ),
          ),
          const SizedBox(height: 20),
        ],
      ),
    ),
  );
}
