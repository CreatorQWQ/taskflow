class TaskModel {
  final int id;
  final String title;
  final String content;
  final String status;
  final DateTime createdAt;

  TaskModel({
    required this.id,
    required this.title,
    required this.content,
    required this.status,
    required this.createdAt,
  });

  // 这是一个工厂方法，把后端传来的 JSON 变成 Dart 对象
  factory TaskModel.fromJson(Map<String, dynamic> json) {
    return TaskModel(
      id: json['id'],
      title: json['title'] ?? '',
      content: json['content'] ?? '',
      status: json['status'] ?? 'pending',
      // 解析后端的 ISO 时间字符串
      createdAt: DateTime.parse(json['created_at']),
    );
  }
}